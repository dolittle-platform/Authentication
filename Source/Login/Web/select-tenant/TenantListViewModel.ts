// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { injectable } from 'tsyringe';

import { IViewContext } from '../MVVM/IViewContext';

import { TenantListProps } from './TenantList';

@injectable()
export class TenantListViewModel {
    async activate({props}: IViewContext<TenantListViewModel, TenantListProps>) {
        props.loading();
        await this.fetchFlow(props.flow);
        props.loaded();
    }
    
    async fetchFlow(flowId: string | null): Promise<void> {
        // TODO: Error handling
        const flow = await (await fetch(`/.auth/self-service/tenant/flows?login_challenge=${flowId}`)).json();
        console.log(flow)

    }
}
