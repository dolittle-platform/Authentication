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

    async activate({props}: IViewContext<TenantListViewModel, TenantListProps>) {
        props.loading();
        await this.fetchFlow(props.flow);
        props.loaded();
    }

    /*
     "{
        "ID":"5b03d0ea800247cf8a8f36a5729eac63",
        "AvailableTenants":[
            {
                "ID":"cc1179a4-d077-45da-a29f-4dbc1cd43e7d",
                "Display":"cc1179a4-d077-45da-a29f-4dbc1cd43e7d"
            },
            {
                "ID":"jakobs-tenant",
                "Display":"jakobs-tenant"
            }
        ]
    }"
    */
    async fetchFlow(flowId: string | null): Promise<void> {
        // TODO: Error handling
        const flow = await (await fetch(`/.auth/self-service/tenant/flows?login_challenge=${flowId}`)).json();
        console.log(flow)
        this.id = flow.ID;
        this.tenants = flow.User.Tenants.map(tenant => ({ id: tenant.ID, display: tenant.Display}))
    }
}
