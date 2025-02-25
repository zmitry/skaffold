/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcb

import (
	"context"
	"fmt"

	"google.golang.org/api/cloudbuild/v1"

	latestV1 "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest/v1"
)

func (b *Builder) koBuildSpec(ctx context.Context, a *latestV1.Artifact, tag string) (cloudbuild.Build, error) {
	k := a.KoArtifact

	return cloudbuild.Build{
		Steps: []*cloudbuild.BuildStep{{
			Name: "gcr.io/es-scalability-test/ko",
			Args: []string{fmt.Sprintf("%s", k.Dir)},
			Env:  k.Env,
		}},
	}, nil
}
