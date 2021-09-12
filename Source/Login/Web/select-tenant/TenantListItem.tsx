// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import { Tenant } from './Tenant';

export type TenantListItemProps = {
    tenant: Tenant;
};

const useStyles = makeStyles({
    button: {
        width: '250px',
    },
});

export const TenantListItem = (props: TenantListItemProps): JSX.Element => {
    const classes = useStyles();
    return (
        <Button
            variant="contained"
            className={classes.button}
            name="tenant"
            value={props.tenant.id}
            type="submit"
        >{props.tenant.display}</Button>
    );
};
