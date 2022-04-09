// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { FormDescriptor, isFormDescriptor } from '../forms/FormDescriptor';
import { IdentityProvider, isIdentityProvider } from './IdentityProvider';

export type Flow = {
    id: string;
    refresh: boolean;
    form: FormDescriptor;
    providers: IdentityProvider[];
};

export const isFlow = (obj: any): obj is Flow => {
    if (typeof obj.id !== 'string') return false;
    if (typeof obj.refresh !== 'boolean') return false;

    if (typeof obj.form !== 'object') return false;
    if (!isFormDescriptor(obj.form)) return false;

    if (!Array.isArray(obj.providers)) return false;
    for (const provider of obj.providers) {
        if (!isIdentityProvider(provider)) {
            return false;
        }
    }

    return true;
};
