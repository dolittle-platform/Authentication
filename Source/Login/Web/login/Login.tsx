// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';
import { useLocation } from 'react-router';

import { LoginProviderList } from './LoginProviderList'

import './Login.scss';
import { LoginViewModel } from './LoginViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export type LoginProps = {
    loading: Function;
    loaded: Function;
};

export const Login = withViewModel<LoginViewModel, LoginProps>(LoginViewModel, ({ viewModel, props }) => {
    return (
        <>
            <h2>Select login provider:</h2>
            <LoginProviderList loading={props.loading} loaded={props.loaded} challenge={getChallenge()} />
        </>
    );
});

const getChallenge = () => {
    const query = new URLSearchParams(useLocation().search);
    return query.get('challenge');
}