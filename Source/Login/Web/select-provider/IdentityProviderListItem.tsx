// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    provider: IdentityProvider;
};

export const IdentityProviderListItem = ({ provider }: IdentityProviderListItemProps): JSX.Element => {
    const { id, imageURL, display } = provider;
    return (
        <Button
            variant='outlined'
            name='provider'
            value={id}
            type='submit'
            startIcon={
                <img
                    src={imageURL}
                    style={{ maxInlineSize: '20px', maxBlockSize: '20px' }}
                />
            }
        >
            Sign in with {display}
        </Button>
    );
}
