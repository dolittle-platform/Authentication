// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box } from '@mui/material';

import { Tenant } from './Tenant';
import { TenantListItem } from './TenantListItem';

export type TenantListProps = {
    tenants: Tenant[];
};

export const TenantList = ({ tenants }: TenantListProps) =>
    <Box sx={{ display: 'inline-block' }}>
        {tenants.map(tenant => (
            <TenantListItem key={tenant.id} tenant={tenant} />
        ))}
    </Box>;
