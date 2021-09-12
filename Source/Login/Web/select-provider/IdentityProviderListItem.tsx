// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    provider: IdentityProvider;
};

const useStyles = makeStyles({
    icon: {
        maxWidth: '20px',
        maxHeight: '20px',
    },
    button: {
        width: '250px',
    },
});

export const IdentityProviderListItem = (props: IdentityProviderListItemProps): JSX.Element => {
    const classes = useStyles();
    return (
        <Button
            variant="contained"
            className={classes.button}
            name="provider"
            value={props.provider.id}
            type="submit"
            startIcon={<img src={props.provider.imageURL} className={classes.icon} />}
        >Sign in with {props.provider.display}</Button>
    );
}
