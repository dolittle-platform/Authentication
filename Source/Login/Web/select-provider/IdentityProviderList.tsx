// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { IdentityProviderListViewModel } from './IdentityProviderListViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export type IdentityProviderListProps = {
    loading: Function;
    loaded: Function;
    flow: string | null;
};

export const IdentityProviderList = withViewModel<IdentityProviderListViewModel, IdentityProviderListProps>(IdentityProviderListViewModel, ({ viewModel, props }) => {
    return (
        <>
            {viewModel.providers.map((provider => {
                return <h3 key={provider.name}>{provider.name}</h3>;
            }))}
        </>
    );
});
