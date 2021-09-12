// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { InvalidIdentityProviderProperty } from './InvalidIdentityProviderProperty';

export type IdentityProvider = {
    id: string;
    display: string;
    imageURL: string;
};

export const asIdentityProvider = (obj: any): IdentityProvider => {
    if (typeof obj.ID !== 'string') throw new InvalidIdentityProviderProperty('ID');
    if (typeof obj.Display !== 'string') throw new InvalidIdentityProviderProperty('Display');
    if (typeof obj.ImageURL !== 'string') throw new InvalidIdentityProviderProperty('ImageURL');

    return {
        id: obj.ID,
        display: obj.Display,
        imageURL: obj.ImageURL,
    };
};
