---
title: "Shape Up: The Pitfalls They Don't Talk About"
date: "2025-06-11"
tags: ["product", "people", "organization"]
excerpt: "Shape Up is a compelling product development methodology, but it comes with hidden challenges that teams need to understand before adopting it."
---

First and foremost, I am not downplaying [Shape Up](https://basecamp.com/shapeup) and its importance to IT professionals, nor am I recommending against it. This is simply 1:13am blabber about a few thoughts that I have that I think are under-emphasized, or perhaps brushed under the rug. So anyway, here is my

# **(☞⌐▀͡ ͜ʖ͡▀ )☞ Top 4 Pitfalls for those working in, or thinking about working with, Shape Up (☞⌐▀͡ ͜ʖ͡▀ )☞**

## Cooling Down Correctly

The most challenging aspect of Shape Up is accepting that **cool-downs might be lost productivity**. The Shape Up explicitly encourages engineers to use cool-downs for innovation, tech debt reduction, and exploration. While this sounds fantastic in theory, the reality is messier.

Engineers will genuinely try their best during these periods — they'll experiment with new technologies, refactor legacy code, and explore creative solutions. But sometimes these efforts don't pan out. That shiny new framework doesn't solve the problem you thought it would. The refactoring uncovers more complexity than expected. The innovative feature prototype reveals fundamental flaws.

This is not a failure of the engineers — it's an expected outcome of the methodology. Shape Up requires leadership and stakeholders to be genuinely okay with this apparent "lost productivity." If you're measuring success purely by feature delivery velocity, cool-downs will feel wasteful. The value they provide is often intangible and long-term, making them difficult to justify in the short term.

## Everything Must Be Halt-able

Shape Up's betting approach requires a fundamental shift in how we think about commitments. **Every item in a cycle must be halt-able without catastrophic consequences.** This means that if an appetite is exceeded, you must be prepared to stop work, re-evaluate the complexity, and potentially re-bet with new knowledge.

The discipline required to halt work and re-bet is often underestimated. It requires strong leadership support, a culture that views stopping as a sign of good judgment rather than failure, and crucially, **the freedom to fail**. Teams need to be comfortable with the idea that sometimes the best decision is to step back and reshape the problem—or even abandon it entirely if the complexity doesn't justify the value. This freedom to fail is essential for Shape Up to work, but it's often the hardest cultural shift for organizations to embrace.

## The Betting Table Bias

Another overlooked pitfall of Shape Up is how **the composition of your betting table indirectly drives product direction**. The methodology emphasizes the importance of having the right people at the betting table, but it doesn't adequately address how different perspectives can create systematic biases in what gets prioritized.

If your betting meetings are dominated by marketing stakeholders, you'll inevitably end up with a highly marketable product that sounds amazing in demos but may have terrible user experience. Marketing-driven bets tend to prioritize features that look impressive in screenshots and sales presentations, but don't necessarily solve real user problems effectively.

Conversely, if support staff have a strong voice at the betting table, you might find yourself constantly prioritizing customer-specific fixes that address individual complaints but don't solve larger, systemic problems. Each bet feels justified—after all, you're helping customers—but you're essentially playing whack-a-mole instead of addressing root causes.

The same bias applies to other stakeholders: sales teams will push for features that help close deals, engineers might over-prioritize technical debt, and executives could favor high-visibility projects that look good in board presentations. **The betting table becomes a microcosm of organizational politics**, and whatever voice is loudest or most persuasive will shape your product roadmap.

There isn't a right or wrong answer when it comes to how to organize a betting table, but the best approach is to be aware of the bias.

## The Deadline Dilemma

Deadlines are great! Oh wait... you said due tomorrow?

Shape Up is that **items with deadlines or non-negotiable work don't fit the methodology.** Incidents need immediate attention. Critical bugs can't wait for the next betting cycle. Urgent customer requests often have immovable deadlines tied to business commitments.

## Final Thoughts

Honestly, as much as I love providing answers, I don't have a definitive answer—or perhaps there really isn't one. It's just a list of compromises you're willing to make compared to whatever other software development methodology you're considering.

I do have _some_ fleeting thoughts that may be beneficial for anyone interested in Shape Up:

- **Let Senior Engineers drive cool-down work.** This will decrease the potential wastage in efforts that don't have any legs to begin with. Yes, CQRS might work well but currently our `/do/the/thing` endpoint has slowed down by 50% for no apparent reason!
- **Set expectations of quality.** ~~Engineers~~ People, often have the best intentions in mind. Juggling quality, speed, and delivery is complex. Engineers have to feel OK with cutting corners at times, knowing that it will be paid back at some point. Conversely, some definition of quality must be set. The default mindset should be "in X language/framework/library, it is recommended to do it Y way." Unless there truly is a constraint on quality, cutting corners should be a last resort.
- **Testing must be included in the appetite.** However this may manifest in your team structure, if the feature is complete just at the cusp of consuming the appetite, then it is simply not done and should likely be halted. It should be OK to say "we've paused this as too many unforeseen issues cropped up."
- **Your developer and testing workflow should be unified.** I don't have an answer on how to do this, but appetite fundamentally does not work if these two components aren't in unison. If your org only has Software Engineers and you validate through metrics and canaries, then this should be seamless. Alternatively, if you have a traditional Software Engineer/QA Engineer setup, then the interaction should not be disruptive.

Anyway, this is the part where I put final thoughts but I don't really have any right now. If you've read the entire thing, thanks. If you're just skipping to the end, then that's cool too. Right now I'm getting back into some of the old bands I used to listen to as a teenager, like Slipknot and My Chemical Romance. Where is this paragraph going? Who knows

cya
