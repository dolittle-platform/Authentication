// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React, { useEffect, useState } from 'react';
import { Spinner, SpinnerSize } from 'office-ui-fabric-react';
import { Route, Switch, useLocation } from 'react-router-dom';

import { SelectProvider } from '../select-provider/SelectProvider';

import './Layout.scss';
import { LayoutViewModel } from './LayoutViewModel';
import { withViewModel } from '../MVVM/withViewModel';

export const Layout = withViewModel(LayoutViewModel, ({ viewModel }) => {
    const [loadingSpinner, setLoadingSpinner] = useState(true);
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
                    </Switch>
                </div>
            </div>
        </div>
    );
});
