// Copyright © 2021 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package members

import (
	"context"
	"fmt"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
)

func process(ctx context.Context, data *dataIn) (*dataOut, error) {
	if data == nil {
		return nil, errors.New("no data")
	}

	if data.epoch < data.chainTime.AltairInitialEpoch() {
		return nil, errors.New("not an Altair epoch")
	}

	syncCommittee, err := data.eth2Client.(eth2client.SyncCommitteesProvider).SyncCommittee(ctx, fmt.Sprintf("%d", data.chainTime.FirstSlotOfEpoch(data.epoch)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain sync committee information")
	}

	if syncCommittee == nil {
		return nil, errors.New("no sync committee returned")
	}

	results := &dataOut{
		debug:      data.debug,
		quiet:      data.quiet,
		verbose:    data.verbose,
		validators: syncCommittee.Validators,
	}

	return results, nil
}
