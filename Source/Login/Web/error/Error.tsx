// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import './Error.scss';

import { ErrorViewModel } from './ErrorViewModel';
import { withViewModel } from '../MVVM/withViewModel';


export type ErrorProps = {
};

export const Error = withViewModel<ErrorViewModel, ErrorProps>(ErrorViewModel, ({ viewModel, props }) => {
    return (
        <div>
            <h1>Oops, something went wrong...</h1>
        </div>
    );
});