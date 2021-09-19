// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { isTenant, Tenant } from './Tenant';

export type User = {
    subject: string;
    tenants: Tenant[];
};

export const isUser = (obj: any): obj is User => {
    if (typeof obj.subject !== 'string') return false;
    
    if (!Array.isArray(obj.tenants)) return false;
    for (const tenant of obj.tenants) {
        if (!isTenant(tenant)) {
            return false;
        }
    }

    return true;
};
