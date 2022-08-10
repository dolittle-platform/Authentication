// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';

import { Tenant } from './Tenant';

export type TenantListItemProps = {
    tenant: Tenant;
};

export const TenantListItem = ({ tenant: { id, display} }: TenantListItemProps): JSX.Element => (
    <Button
        variant='contained'
        name='tenant'
        value={id}
        type='submit'
        sx={{ minWidth: '9.375rem' }}
    >{display}</Button>
);
