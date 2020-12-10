// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React, { useState } from 'react';
import { Spinner, SpinnerSize } from 'office-ui-fabric-react';
import { Route, Switch } from 'react-router-dom';

import { SelectProvider } from '../select-provider/SelectProvider';
import { SelectTenant } from '../select-tenant/SelectTenant';
import { Error } from '../error/Error';

import './Layout.scss';
import { LayoutViewModel } from './LayoutViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export const Layout = withViewModel(LayoutViewModel, ({ viewModel }) => {
    const [loadingSpinner, setLoadingSpinner] = useState(false);
    const contentLoading = () => setLoadingSpinner(true);
    const contentLoaded = () => setLoadingSpinner(false);

    return (
        <div className='application'>
            <div className="main">
                <div className="content">
                    <div className="spinner">
                        <Spinner styles={{ root: { display: loadingSpinner ? undefined! : 'none' } }} size={SpinnerSize.large} label="Loading Content" />
                    </div>

                    <Switch>
                        <Route path="/.auth/select-provider">
                            <SelectProvider loading={contentLoading} loaded={contentLoaded} />
                        </Route>
                        <Route path="/.auth/select-tenant">
                            <SelectTenant loading={contentLoading} loaded={contentLoaded} />
                        </Route>
                        <Route path="/.auth/error">
                            <Error />
                        </Route>
                    </Switch>
                </div>
            </div>
        </div>
    );
});
