// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Endpoint } from '@rest-hooks/endpoint';

import { Flow, asFlow } from './Flow';

const fetchFlow = async (flowID: string | null): Promise<Flow | null> => {
    if (flowID === null) return null;

    const response = await fetch(`/.auth/self-service/login/flows?id=${flowID}`);
    const data = await response.json();

    return asFlow(data);
};

export const Flows = new Endpoint(fetchFlow);
