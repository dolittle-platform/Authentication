// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export type Tenant = {
    id: string;
    display: string;
};

export const isTenant = (obj: any): obj is Tenant => {
    if (typeof obj.id !== 'string') return false;
    if (typeof obj.display !== 'string') return false;
    return true;
};
