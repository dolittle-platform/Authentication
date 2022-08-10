// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    provider: IdentityProvider;
};

export const IdentityProviderListItem = ({ provider: { id, imageURL, display } }: IdentityProviderListItemProps): JSX.Element => (
        <Button
            variant='outlined'
            name='provider'
            value={id}
            type='submit'
            startIcon={
                <img
                    src={imageURL}
                    style={{ maxWidth: '1.25rem', maxHeight: '1.25rem' }}
                />
            }
        >
            Sign in with {display}
        </Button>
);
