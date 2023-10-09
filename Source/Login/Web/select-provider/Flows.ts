// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Endpoint } from '@rest-hooks/endpoint';

import { Flow, isFlow } from './Flow';
import { InvalidFlow } from './InvalidFlow';

const fetchFlow = async (flowID: string | null): Promise<Flow | null> => {
    if (flowID === null) return null;

    const response = await fetch(`/.auth/self-service/login/flows?id=${flowID}`);
    const data = await response.json();

    if (!isFlow(data)) {
        throw new InvalidFlow(data);
    }

    return data;
};

export const Flows = new Endpoint(fetchFlow);
