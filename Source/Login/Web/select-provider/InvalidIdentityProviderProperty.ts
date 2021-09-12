// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export class InvalidIdentityProviderProperty extends Error {
    constructor(property: string) {
        super(`The login identity provider was not valid because the field '${property}' was missing or invalid.`);
    }
};
