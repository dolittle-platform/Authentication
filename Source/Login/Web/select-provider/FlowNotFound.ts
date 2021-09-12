// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export class FlowNotFound extends Error {
    constructor(id: string | null) {
        super(`The login flow was not found for id '${id}'.`);
    }
};
