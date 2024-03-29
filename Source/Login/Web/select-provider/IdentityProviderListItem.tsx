// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Button } from '@mui/material';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    provider: IdentityProvider;
};

export const IdentityProviderListItem = ({ provider: { id, imageURL, display } }: IdentityProviderListItemProps) =>
    <Button
        variant='outlined'
        name='provider'
        value={id}
        type='submit'
        startIcon={
            <img src={imageURL} style={{ maxWidth: '1.25rem', maxHeight: '1.25rem' }} />
        }
        sx={{ display: 'flex', width: 1 }}
    >
        Sign in with {display}
    </Button>;
