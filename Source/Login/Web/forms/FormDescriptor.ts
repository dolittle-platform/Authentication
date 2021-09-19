// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.


export type FormSubmitMethod = 'GET' | 'POST';

export type FormDescriptor = {
    submitMethod: FormSubmitMethod;
    submitAction: string;
    csrfToken: string | undefined;
};

export const isFormDescriptor = (obj: any): obj is FormDescriptor => {
    if (obj.submitMethod !== 'GET' && obj.submitMethod !== 'POST') return false;
    if (typeof obj.submitAction !== 'string') return false;
    if (typeof obj.csrfToken !== 'string' && typeof obj.csrfToken !== 'undefined') return false;
    return true;
};
