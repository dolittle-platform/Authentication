package login

import (
	"errors"
	"net/url"

	"dolittle.io/login/flows/forms"
	"dolittle.io/login/providers"
	ory "github.com/ory/kratos-client-go"
)

type Parser interface {
	ParseLoginFlowFrom(response *ory.SelfServiceLoginFlow) (*Flow, error)
}

func NewParser(configuration Configuration, providers providers.Getter) Parser {
	return &parser{
		configuration: configuration,
		providers:     providers,
	}
}

type parser struct {
	configuration Configuration
	providers     providers.Getter
}

func (p *parser) ParseLoginFlowFrom(response *ory.SelfServiceLoginFlow) (*Flow, error) {

	flowID, ok := response.GetIdOk()
	if !ok {
		return nil, errors.New("flow id not set")
	}

	forced, ok := response.GetForcedOk()
	if !ok {
		return nil, errors.New("flow id not set")
	}

	formUI, ok := response.GetUiOk()
	if !ok {
		return nil, errors.New("flow form not set")
	}
	uiNodes, ok := formUI.GetNodesOk()
	if !ok {
		return nil, errors.New("form nodes not set")
	}

	csrfToken, err := p.getCSRFToken(uiNodes)
	if err != nil {
		return nil, err
	}

	submitAction, ok := formUI.GetActionOk()
	if !ok {
		return nil, errors.New("form submit action not set")
	}
	_, err = url.Parse(*submitAction)
	if err != nil {
		return nil, err
	}

	submitMethod, ok := formUI.GetMethodOk()
	if !ok {
		return nil, errors.New("form submit method not set")
	}

	providerIDs, err := p.getProviderIDs(uiNodes)
	if err != nil {
		return nil, err
	}

	providers := make([]providers.Provider, 0)
	for _, providerID := range providerIDs {
		provider, err := p.providers.GetProviderByID(providerID)
		if err != nil {
			return nil, err
		}
		providers = append(providers, provider)
	}

	return &Flow{
		ID:     FlowID(*flowID),
		Forced: *forced,
		Form: forms.Form{
			SubmitMethod: *submitMethod,
			SubmitAction: *submitAction,
			CSRFToken:    csrfToken,
		},
		Providers: providers,
	}, nil
}

func (p *parser) getCSRFToken(nodes []ory.UiNode) (string, error) {
	tokens, err := p.getFormFieldStringValues(p.configuration.CSRFTokenFieldName(), nodes)
	if err != nil {
		return "", err
	}
	if len(tokens) > 1 {
		return "", errors.New("more than one csrf token found")
	}
	if len(tokens) < 1 {
		return "", errors.New("no csrf tokens found")
	}
	return tokens[0], nil
}

func (p *parser) getProviderIDs(nodes []ory.UiNode) ([]string, error) {
	return p.getFormFieldStringValues(p.configuration.ProviderFieldName(), nodes)
}

func (p *parser) getFormFieldStringValues(fieldName string, nodes []ory.UiNode) ([]string, error) {
	values := []string{}
	for _, node := range nodes {
		attributes, ok := node.GetAttributesOk()
		if !ok {
			continue
		}
		if name, ok := attributes.UiNodeInputAttributes.GetNameOk(); !ok || *name != fieldName {
			continue
		}
		value, ok := attributes.UiNodeInputAttributes.GetValueOk()
		if !ok {
			return nil, errors.New("form value was not set")
		}
		stringValue, ok := (*value).(string)
		if !ok {
			return nil, errors.New("form value was not string")
		}
		values = append(values, stringValue)
	}
	return values, nil
}
