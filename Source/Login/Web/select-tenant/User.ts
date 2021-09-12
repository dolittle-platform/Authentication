// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { InvalidUserProperty } from './InvalidUserProperty';
import { asTenant, Tenant } from './Tenant';

export type User = {
    subject: string;
    tenants: Tenant[];
};

export const asUser = (obj: any): User => {
    if (typeof obj.Subject !== 'string') throw new InvalidUserProperty('Subject');
    if (!Array.isArray(obj.Tenants)) throw new InvalidUserProperty('Tenants');

    return {
        subject: obj.Subject,
        tenants: obj.Tenants.map((obj: any) => asTenant(obj)),
    };
};
