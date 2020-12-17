// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { DocumentCard, DocumentCardDetails, DocumentCardPreview, DocumentCardTitle, DocumentCardType, ImageFit } from 'office-ui-fabric-react';
import React from 'react';

import { IdentityProvider } from './IdentityProvider';

export type IdentityProviderListItemProps = {
    formAction: string;
    formMethod: string;
    formCsrfToken: string;
    provider: IdentityProvider;
};

export class IdentityProviderListItem extends React.Component<IdentityProviderListItemProps> {
    form = React.createRef<HTMLFormElement>();

    onClick = () => {
        this.form.current?.submit()
    };

    render() {
        return (
            <DocumentCard type={DocumentCardType.compact} onClick={this.onClick}>
                { this.props.provider.imageURL &&
                    <DocumentCardPreview previewImages={[
                        {
                            previewImageSrc: this.props.provider.imageURL,
                            width: 60,
                            height: 60,
                        }
                    ]} />
                }
                <DocumentCardDetails>
                    <DocumentCardTitle title={this.props.provider.display} />
                    <form ref={this.form} action={this.props.formAction} method={this.props.formMethod}>
                        <input type="hidden" name="csrf_token" value={this.props.formCsrfToken} />
                        <input type="hidden" name="provider" value={this.props.provider.id} />
                    </form>
                </DocumentCardDetails>
            </DocumentCard>
        );
    }
}