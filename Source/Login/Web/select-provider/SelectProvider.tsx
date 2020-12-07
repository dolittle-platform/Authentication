// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';
import { useLocation } from 'react-router';

import { IdentityProviderList } from './IdentityProviderList';

import './SelectProvider.scss';
import { SelectProviderViewModel } from './SelectProviderViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export type SelectProviderProps = {
    loading: Function;
    loaded: Function;
};

export const SelectProvider = withViewModel<SelectProviderViewModel, SelectProviderProps>(SelectProviderViewModel, ({ viewModel, props }) => {
    return (
        <>
            <h2>Select login provider:</h2>
            <IdentityProviderList loaded={props.loaded} loading={props.loading} flow={getFlow()} />
        </>
    );
});

const getFlow = () => {
    const query = new URLSearchParams(useLocation().search);
    return query.get('flow');
}
