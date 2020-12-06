// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { LoginProviderListViewModel } from './LoginProviderListViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export type LoginProviderListProps = {
    loading: Function;
    loaded: Function;
    challenge: string | null;
};

export const LoginProviderList = withViewModel<LoginProviderListViewModel, LoginProviderListProps>(LoginProviderListViewModel, ({ viewModel, props }) => {
    return (
        <>
            {viewModel.providers.map((provider => {
                return <h3 key={provider.name}>{provider.name}</h3>;
            }))}
        </>
    );
});