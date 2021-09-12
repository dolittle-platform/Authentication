// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export class ChallengeNotFound extends Error {
    constructor(id: string | null) {
        super(`The login challenge was not found for id '${id}'.`);
    }
};
