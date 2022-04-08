// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import Box from '@mui/material/Box';

import { IdentityProvider } from './IdentityProvider';
import { IdentityProviderListItem } from './IdentityProviderListItem';

export type IdentityProviderListProps = {
    providers: IdentityProvider[];
};

export const IdentityProviderList = (props: IdentityProviderListProps): JSX.Element => {
    return (
        <>
            {
                props.providers.map(provider => (
                    <Box key={provider.id} css={{ marginBottom: '28px' }}>
                        <IdentityProviderListItem provider={provider} />
                    </Box>
                ))
            }
        </>
    );
};
