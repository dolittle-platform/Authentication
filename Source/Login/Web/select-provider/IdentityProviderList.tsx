// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { IdentityProvider } from './IdentityProvider';
import { IdentityProviderListItem } from './IdentityProviderListItem';

export type IdentityProviderListProps = {
    providers: IdentityProvider[];
};

export const IdentityProviderList = ({ providers }: IdentityProviderListProps): JSX.Element => {
    return (
        <>
            {
                providers.map(provider => (
                    <IdentityProviderListItem key={provider.id} provider={provider} />
                ))
            }
        </>
    );
};
