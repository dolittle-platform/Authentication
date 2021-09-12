// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';

import { IdentityProvider } from './IdentityProvider';
import { IdentityProviderListItem } from './IdentityProviderListItem';

export type IdentityProviderListProps = {
    providers: IdentityProvider[];
};

const useStyles = makeStyles({
    row: {
        marginBottom: '28px',
    },
});

export const IdentityProviderList = (props: IdentityProviderListProps): JSX.Element => {
    const classes = useStyles();
    return (
        <>
            {
                props.providers.map(provider => (
                    <Box key={provider.id} className={classes.row}>
                        <IdentityProviderListItem provider={provider} />
                    </Box>
                ))
            }
        </>
    );
};
