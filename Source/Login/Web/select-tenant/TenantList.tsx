// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';

import { Tenant } from './Tenant';
import { TenantListItem } from './TenantListItem';

export type TenantListProps = {
    tenants: Tenant[];
};

export const TenantList = (props: TenantListProps): JSX.Element => {
    return (
        <>
            {
                props.tenants.map(tenant => (
                    <Box key={tenant.id} mb={2}>
                        <TenantListItem tenant={tenant} />
                    </Box>
                ))
            }
        </>
    );
};
