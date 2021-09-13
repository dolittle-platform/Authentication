// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

declare global {
    interface Window {
        configuration: Configuration | undefined
    }
}

export type Configuration = {
    showDolittleHeadline: boolean | undefined,
    applicationName: string | undefined,
    supportEmail: string | undefined,
};

export const configuration: Configuration = window.configuration ?? {
    showDolittleHeadline: true,
    applicationName: 'Dolittle Studio',
    supportEmail: 'support@dolittle.com',
};
