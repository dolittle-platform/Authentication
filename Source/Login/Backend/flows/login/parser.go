package login

import (
	"errors"
	"net/url"

	"dolittle.io/login/providers"
	"github.com/ory/kratos-client-go/models"
)

type Parser interface {
	ParseLoginFlowFrom(response *models.LoginFlow) (*Flow, error)
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

func (p *parser) ParseLoginFlowFrom(response *models.LoginFlow) (*Flow, error) {
	method, ok := response.Methods["oidc"]
	if !ok {
		return nil, errors.New("response does not have oidc method")
	}

	if method.Config.Action == nil {
		return nil, errors.New("form submit action not set")
	}
	submitAction, err := url.Parse(*method.Config.Action)
	if err != nil {
		return nil, err
	}

	if method.Config.Method == nil {
		return nil, errors.New("form submit method not set")
	}

	csrfToken, err := p.getCSRFToken(method.Config.Fields)
	if err != nil {
		return nil, err
	}

	providerIDs, err := p.getProviderIDs(method.Config.Fields)
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
		ID:               FlowID(response.ID),
		Forced:           response.Forced,
		FormCSRFToken:    csrfToken,
		FormSubmitAction: submitAction,
		FormSubmitMethod: *method.Config.Method,
		Providers:        providers,
	}, nil
}

func (p *parser) getCSRFToken(fields models.FormFields) (string, error) {
	tokens, err := p.getFormFieldStringValues(p.configuration.CSRFTokenFieldName(), fields)
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

func (p *parser) getProviderIDs(fields models.FormFields) ([]string, error) {
	return p.getFormFieldStringValues(p.configuration.ProviderFieldName(), fields)
}

func (p *parser) getFormFieldStringValues(fieldName string, fields models.FormFields) ([]string, error) {
	values := []string{}
	for _, field := range fields {
		if *field.Name != fieldName {
			continue
		}
		value, isString := field.Value.(string)
		if !isString {
			return nil, errors.New("form value was not string")
		}
		values = append(values, value)
	}
	return values, nil
}
