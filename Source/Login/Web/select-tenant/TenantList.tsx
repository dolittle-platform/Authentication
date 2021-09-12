// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';

import { Tenant } from './Tenant';
import { TenantListItem } from './TenantListItem';
// import { TenantListItem } from './TenantListItem';

export type TenantListProps = {
    tenants: Tenant[];
};

const useStyles = makeStyles({
    row: {
        marginBottom: '28px',
    },
});

export const TenantList = (props: TenantListProps): JSX.Element => {const classes = useStyles();
    return (
        <>
            {
                props.tenants.map(tenant => (
                    <Box key={tenant.id} className={classes.row}>
                        <TenantListItem tenant={tenant} />
                    </Box>
                ))
            }
        </>
    );
};
