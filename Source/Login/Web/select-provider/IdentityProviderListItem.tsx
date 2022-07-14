// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Button from '@mui/material/Button';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    provider: IdentityProvider;
};

export const IdentityProviderListItem = (props: IdentityProviderListItemProps): JSX.Element => {
    return (
        <Button
            variant='outlined'
            name='provider'
            value={props.provider.id}
            type='submit'
            startIcon={
                <img
                    src={props.provider.imageURL}
                    css={{ maxInlineSize: '20px', maxBlockSize: '20px' }}
                />
            }
        >
            Sign in with {props.provider.display}
        </Button>
    );
}
