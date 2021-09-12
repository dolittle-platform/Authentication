// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { FormDescriptor, asFormDescriptor } from '../forms/FormDescriptor';
import { asIdentityProvider, IdentityProvider } from './IdentityProvider';
import { InvalidFlowProperty } from './InvalidFlowProperty';

export type Flow = {
    id: string;
    forced: boolean;
    formCSRFToken: string;
    providers: IdentityProvider[];
} & FormDescriptor;

export const asFlow = (obj: any): Flow => {
    if (typeof obj.ID !== 'string') throw new InvalidFlowProperty('ID');
    if (typeof obj.Forced !== 'boolean') throw new InvalidFlowProperty('Forced');
    if (typeof obj.FormCSRFToken !== 'string') throw new InvalidFlowProperty('FormCSRFToken');
    if (!Array.isArray(obj.Providers)) throw new InvalidFlowProperty('Providers');

    const formDescriptor = asFormDescriptor(obj);

    return {
        id: obj.ID,
        forced: obj.Forced,
        formCSRFToken: obj.FormCSRFToken,
        providers: obj.Providers.map((obj: any) => asIdentityProvider(obj)),
        ...formDescriptor,
    };
};
