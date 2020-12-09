// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { injectable } from 'tsyringe';

import { IViewContext } from '../MVVM/IViewContext';
import { Tenant } from './Tenant';

import { TenantListProps } from './TenantList';

@injectable()
export class TenantListViewModel {
    id: string = '';
    tenants: Tenant[] = [];
    formAction: string = '';
    formMethod: string = '';

    async activate({props}: IViewContext<TenantListViewModel, TenantListProps>) {
        props.loading();
        await this.fetchFlow(props.flow);
        props.loaded();
    }

    async fetchFlow(flowId: string | null): Promise<void> {
        // TODO: Error handling
        const flow = await (await fetch(`/.auth/self-service/tenant/flows?login_challenge=${flowId}`)).json();
        console.log(flow)
        this.id = flow.ID;
        this.formAction = `${flow.FormSubmitAction.Scheme}://${flow.FormSubmitAction.Host}${flow.FormSubmitAction.Path}`;;
        this.formMethod = flow.FormSubmitMethod;
        this.tenants = flow.User.Tenants.map(tenant => ({ id: tenant.ID, display: tenant.Display}))
    }
}
