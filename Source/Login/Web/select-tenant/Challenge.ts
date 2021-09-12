// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { asFormDescriptor, FormDescriptor } from '../forms/FormDescriptor';
import { InvalidChallengeProperty } from './InvalidChallengeProperty';
import { asUser, User } from './User';

export type Challenge = {
    id: string;
    user: User;
} & FormDescriptor;

export const asChallenge = (obj: any): Challenge => {
    if (typeof obj.ID !== 'string') throw new InvalidChallengeProperty('ID');
    if (typeof obj.User !== 'object') throw new InvalidChallengeProperty('User');

    const user = asUser(obj.User);
    const formDescriptor = asFormDescriptor(obj);

    return {
        id: obj.ID,
        user,
        ...formDescriptor,
    };
};
