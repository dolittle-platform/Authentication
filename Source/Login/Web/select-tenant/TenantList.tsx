// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { withViewModel } from '../MVVM/withViewModel';

import { TenantListViewModel } from './TenantListViewModel';

export type TenantListProps = {
    loading: Function;
    loaded: Function;
    flow: string | null;
};

export const TenantList = withViewModel<TenantListViewModel, TenantListProps>(TenantListViewModel, ({ viewModel, props }) => {
    return (
        <>
            <h1>Select tenant:</h1>
        </>
    );
});
