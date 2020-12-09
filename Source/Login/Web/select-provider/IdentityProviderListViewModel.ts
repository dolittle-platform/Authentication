// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { injectable } from 'tsyringe';

import { IViewContext } from '../MVVM/IViewContext';

import { IdentityProvider } from './IdentityProvider';
import { IdentityProviderListProps } from './IdentityProviderList';

@injectable()
export class IdentityProviderListViewModel {
    flowId: string = '';
    providers: IdentityProvider[] = [];
    forced: boolean = false;
    formCsrfToken: string = '';
    formAction: string = '';
    formMethod: string = '';

    async activate({props}: IViewContext<IdentityProviderListViewModel, IdentityProviderListProps>) {
        props.loading();
        await this.fetchFlow(props.flow);
        props.loaded();
    }
    
    async fetchFlow(flowId: string | null): Promise<void> {
        // TODO: Error handling
        const flow = await (await fetch(`/.auth/self-service/login/flows?id=${flowId}`)).json();
        console.log(flow)

        this.flowId = flow.ID;
        this.forced = flow.Forced;
        this.formCsrfToken = flow.FormCSRFToken;
        this.formAction = `${flow.FormSubmitAction.Scheme}://${flow.FormSubmitAction.Host}${flow.FormSubmitAction.Path}`;
        this.formMethod = flow.FormSubmitMethod;
        this.providers = flow.Providers.map(provider => ({ id: provider.ID, display: provider.Display }));
    }
}

/*
"{
    "ID":"e0ac5be5-b317-41a4-8040-6c8b2701d359",
    "Forced":false,
    "FormCSRFToken":"wx+S+DTmDif/iTr2esr+A2PIFCneYQV72kdMUGid4F2qTSqeiVFflh34pHtGKwo6yfVbaQpUmRmzaohQ/7UCWw==",
    
    "FormSubmitAction":{
        "Scheme":"http",
        "Opaque":"",
        "User":null,
        "Host":"local.dolittle.studio:8080",
        "Path":"/.auth/self-service/methods/oidc/auth/e0ac5be5-b317-41a4-8040-6c8b2701d359",
        "RawPath":"","ForceQuery":false,"RawQuery":"","Fragment":"","RawFragment":""
    },
    "FormSubmitMethod":"POST",
    "Providers":[
        {
        "ID":"sample",
        "Display":"sample"
        }
    ]
}"
*/
