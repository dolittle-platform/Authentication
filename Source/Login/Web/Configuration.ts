// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

export type Configuration = {
    showDolittleHeadline: boolean,
    applicationName: string | undefined,
    supportEmail: string | undefined,
    startPath: string,
    logoutPath: string,
};

type PartialConfiguration = {
    [Property in keyof Configuration]: Configuration[Property] | undefined;
};

declare global {
    interface Window {
        configuration: PartialConfiguration | undefined
    }
}

const defaults: Configuration = {
    showDolittleHeadline: true,
    applicationName: 'Dolittle Studio',
    supportEmail: 'support@dolittle.com',
    startPath: '/',
    logoutPath: '/.auth/cookies/logout'
};

export const configuration = Object.fromEntries([
    ...Object.entries(defaults),
    ...Object.entries(window.configuration ?? {}).filter(([_, value]) => value !== undefined),
]) as Configuration;
