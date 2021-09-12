// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { InvalidTenantProperty } from './InvalidTenantProperty';

export type Tenant = {
    id: string;
    display: string;
};

export const asTenant = (obj: any): Tenant => {
    if (typeof obj.ID !== 'string') throw new InvalidTenantProperty('ID');
    if (typeof obj.Display !== 'string') throw new InvalidTenantProperty('Display');

    return {
        id: obj.ID,
        display: obj.Display,
    };
};
