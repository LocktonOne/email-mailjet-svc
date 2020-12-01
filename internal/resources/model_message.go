/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Message struct {
	Key
	Attributes MessageAttributes `json:"attributes"`
}
type MessageResponse struct {
	Data     Message  `json:"data"`
	Included Included `json:"included"`
}

type MessageListResponse struct {
	Data     []Message `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustMessage - returns Message from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustMessage(key Key) *Message {
	var message Message
	if c.tryFindEntry(key, &message) {
		return &message
	}
	return nil
}
