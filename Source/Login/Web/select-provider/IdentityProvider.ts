// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export type IdentityProvider = {
    id: string;
    display: string;
    imageURL: string;
};

export const isIdentityProvider = (obj: any): obj is IdentityProvider => {
    if (typeof obj.id !== 'string') return false;
    if (typeof obj.display !== 'string') return false;
    if (typeof obj.imageURL !== 'string') return false;
    return true;
};
