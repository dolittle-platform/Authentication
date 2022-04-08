// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Navigate } from 'react-router-dom';

export const RedirectToError = (): JSX.Element => {
    return (
        <Navigate to="/.auth/error" />
    );
};
