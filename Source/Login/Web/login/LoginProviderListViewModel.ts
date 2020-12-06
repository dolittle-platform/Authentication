// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { injectable } from 'tsyringe';

import { IViewContext } from '../MVVM/IViewContext';

import { LoginProvider } from './LoginProvider'
import { LoginProviderListProps } from './LoginProviderList';

@injectable()
export class LoginProviderListViewModel {
    providers: LoginProvider[] = [];

    constructor() {}

    async activate({props}: IViewContext<LoginProviderListViewModel, LoginProviderListProps>) {
        props.loading();
        await this.fetchLoginProviders(props.challenge);
        props.loaded();
    }

    fetchLoginProviders(challenge: string | null): Promise<void> {
        return new Promise((resolve) => {
            if (challenge) {
                setTimeout(() => {
                    console.log('Fetching login providers', challenge)
                    this.providers = [
                        { name: 'Microsoft' },
                        { name: 'GitHub' },
                        { name: 'Facebook' },
                    ];
                    resolve();
                }, 1000);
            } else {
                resolve();
            }
        });
    }
}