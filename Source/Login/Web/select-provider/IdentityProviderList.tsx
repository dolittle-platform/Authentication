// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { IdentityProviderListViewModel } from './IdentityProviderListViewModel';
import { withViewModel } from '../MVVM/withViewModel';
import { IdentityProviderListItem } from './IdentityProviderListItem';

export type IdentityProviderListProps = {
    loading: Function;
    loaded: Function;
    flow: string | null;
};

export const IdentityProviderList = withViewModel<IdentityProviderListViewModel, IdentityProviderListProps>(IdentityProviderListViewModel, ({ viewModel, props }) => {
    return (
        <>
            <h1>Log in with:</h1>
            {viewModel.providers.map(provider =>
                <IdentityProviderListItem
                    key={provider.id}
                    provider={provider}
                    formAction={viewModel.formAction}
                    formMethod={viewModel.formMethod}
                    formCsrfToken={viewModel.formCsrfToken} />
            )}
        </>
    );
});
