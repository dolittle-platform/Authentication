// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';

import { Tenant } from './Tenant';

export type TenantListItemProps = {
    tenant: Tenant;
};

export const TenantListItem = (props: TenantListItemProps): JSX.Element => {
    return (
        <Button
            variant="contained"
            name="tenant"
            value={props.tenant.id}
            type="submit"
            sx={{ minInlineSize: '150px' }}
        >{props.tenant.display}</Button>
    );
};
