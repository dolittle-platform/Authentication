// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box } from '@mui/material';

import { Tenant } from './Tenant';
import { TenantListItem } from './TenantListItem';

export type TenantListProps = {
    tenants: Tenant[];
};

export const TenantList = ({ tenants }: TenantListProps): JSX.Element => (
    <>
        {
            tenants.map(tenant => (
                <Box key={tenant.id} sx={{ mb: 2, inlineSize: '100%' }}>
                    <TenantListItem tenant={tenant} />
                </Box>
            ))
        }
    </>
);
