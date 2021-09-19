// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export class InvalidChallenge extends Error {
    constructor(obj: any) {
        super(`The retrieved select tenant challenge is not valid. Received '${JSON.stringify(obj)}'`);
    }
};
