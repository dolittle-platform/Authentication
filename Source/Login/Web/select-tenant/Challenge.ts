// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { FormDescriptor, isFormDescriptor } from '../forms/FormDescriptor';
import { isUser, User } from './User';

export type Challenge = {
    id: string;
    form: FormDescriptor;
    user: User;
};

export const isChallenge = (obj: any): obj is Challenge => {
    if (typeof obj.id !== 'string') return false;

    if (typeof obj.form !== 'object') return false;
    if (!isFormDescriptor(obj.form)) return false;

    if (typeof obj.user !== 'object') return false;
    if (!isUser(obj.user)) return false;

    return true;
};
