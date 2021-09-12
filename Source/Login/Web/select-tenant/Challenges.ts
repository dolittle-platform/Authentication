// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Endpoint } from '@rest-hooks/endpoint';

import { asChallenge, Challenge } from './Challenge';

const fetchChallenge = async (challengeID: string | null): Promise<Challenge | null> => {
    if (challengeID === null) return null;

    const response = await fetch(`/.auth/self-service/tenant/flows?login_challenge=${challengeID}`);
    const data = await response.json();

    return asChallenge(data);
};

export const Challenges = new Endpoint(fetchChallenge);
