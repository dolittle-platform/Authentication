// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { DocumentCard, DocumentCardDetails, DocumentCardPreview, DocumentCardTitle, DocumentCardType, ImageFit } from 'office-ui-fabric-react';
import React from 'react';

import { Tenant } from './Tenant';

export type TenantListItemProps = {
    formAction: string;
    formMethod: string;
    // formCsrfToken: string;
    flowId: string;
    tenant: Tenant;
};

export class TenantListItem extends React.Component<TenantListItemProps> {
    form = React.createRef<HTMLFormElement>();

    onClick = () => {
        this.form.current?.submit()
    };

    render() {
        return (
            <DocumentCard type={DocumentCardType.compact} onClick={this.onClick}>
                <DocumentCardDetails>
                    <DocumentCardTitle title={this.props.tenant.display} />
                    <form ref={this.form} action={this.props.formAction} method={this.props.formMethod}>
                        <input type="hidden" name="login_challenge" value={this.props.flowId} />
                        <input type="hidden" name="tenant" value={this.props.tenant.id} />
                    </form>
                </DocumentCardDetails>
            </DocumentCard>
        );
    }
}
