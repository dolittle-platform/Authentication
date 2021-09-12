// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { InvalidFormProperty } from './InvalidFormProperty';

export type FormSubmitMethod = 'GET' | 'POST';

export type FormSubmitAction = {
    scheme: 'http' | 'https';
    opaque: string;
    user: any;
    host: string;
    path: string;
    rawPath: string;
    forceQuery: boolean;
    rawQuery: string;
    fragment: string;
    rawFragment: string;
};

export type FormDescriptor = {
    formSubmitMethod: FormSubmitMethod;
    formSubmitAction: FormSubmitAction;
};

export const asFormDescriptor = (obj: any): FormDescriptor => {
    if (obj.FormSubmitMethod !== 'GET' && obj.FormSubmitMethod !== 'POST') throw new InvalidFormProperty('FormSubmitMethod');
    if (typeof obj.FormSubmitAction !== 'object') throw new InvalidFormProperty('FormSubmitAction');
    if (obj.FormSubmitAction.Scheme !== 'http' && obj.FormSubmitAction.Scheme !== 'https') throw new InvalidFormProperty('FormSubmitAction.Scheme');
    if (typeof obj.FormSubmitAction.Opaque !== 'string') throw new InvalidFormProperty('FormSubmitAction.Opaque');
    if (typeof obj.FormSubmitAction.Host !== 'string') throw new InvalidFormProperty('FormSubmitAction.Host');
    if (typeof obj.FormSubmitAction.Path !== 'string') throw new InvalidFormProperty('FormSubmitAction.Path');
    if (typeof obj.FormSubmitAction.RawPath !== 'string') throw new InvalidFormProperty('FormSubmitAction.RawPath');
    if (typeof obj.FormSubmitAction.ForceQuery !== 'boolean') throw new InvalidFormProperty('FormSubmitAction.ForceQuery');
    if (typeof obj.FormSubmitAction.RawQuery !== 'string') throw new InvalidFormProperty('FormSubmitAction.RawQuery');
    if (typeof obj.FormSubmitAction.Fragment !== 'string') throw new InvalidFormProperty('FormSubmitAction.Fragment');
    if (typeof obj.FormSubmitAction.RawFragment !== 'string') throw new InvalidFormProperty('FormSubmitAction.RawFragment');

    return {
        formSubmitMethod: obj.FormSubmitMethod,
        formSubmitAction: {
            scheme: obj.FormSubmitAction.Scheme,
            opaque: obj.FormSubmitAction.Opaque,
            user: obj.FormSubmitAction.User,
            host: obj.FormSubmitAction.Host,
            path: obj.FormSubmitAction.Path,
            rawPath: obj.FormSubmitAction.RawPath,
            forceQuery: obj.FormSubmitAction.ForceQuery,
            rawQuery: obj.FormSubmitAction.RawQuery,
            fragment: obj.FormSubmitAction.Fragment,
            rawFragment: obj.FormSubmitAction.RawFragment,
        },
    };
};
