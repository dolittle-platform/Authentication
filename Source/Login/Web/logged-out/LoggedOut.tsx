// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';

import { configuration } from '../Configuration';
import { Theme } from '../styles/Theme';

export const LoggedOut = (): JSX.Element => {
    return (
        <>
            <Typography variant='h2'>You have been logged out.</Typography>
            <Button size='large' color='inherit' href={configuration.startPath} css={{ marginTop: '128px', [Theme.breakpoints.up('sm')]: { marginTop: '25px' } }}>Go back and log in again</Button>
        </>
    );
}
