// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Endpoint } from '@rest-hooks/endpoint';

import { Challenge, isChallenge } from './Challenge';
import { InvalidChallenge } from './InvalidChallenge';

const fetchChallenge = async (challengeID: string | null): Promise<Challenge | null> => {
    if (challengeID === null) return null;

    const response = await fetch(`/.auth/self-service/tenant/flows?login_challenge=${challengeID}`);
    const data = await response.json();

    if (!isChallenge(data)) throw new InvalidChallenge(data);

    return data;
};

export const Challenges = new Endpoint(fetchChallenge);
