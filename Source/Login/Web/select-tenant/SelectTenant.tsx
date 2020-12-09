// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';
import { useLocation } from 'react-router';

import './SelectTenant.scss';
import { SelectTenantViewModel } from './SelectTenantViewModel';
import { withViewModel } from '../MVVM/withViewModel';
import { TenantList } from './TenantList';

export type SelectTenantProps = {
    loading: Function;
    loaded: Function;
};

export const SelectTenant = withViewModel<SelectTenantViewModel, SelectTenantProps>(SelectTenantViewModel, ({ viewModel, props }) => {
    return <TenantList flow={getFlow()} loading={props.loading} loaded={props.loaded} />;
});

const getFlow = () => {
    const query = new URLSearchParams(useLocation().search);
    return query.get('login_challenge');
}
